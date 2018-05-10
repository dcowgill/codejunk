from re import compile

import itertools

# Query/indexing parsing needs symmetric processing
DELIMITERS = '\s+|\.+|,+'

class Term:
    """
        Term represents a word in the index, its character offset from the
        beginning of the original document and its relative word offset.
    """

    def __init__(self, term, char_offset, word_offset):
        self.term = term.lower()
        self.char_offset = char_offset
        self.word_offset = word_offset

    def __str__(self):
        return self.term + "[" + str(self.word_offset) + "]"

class TermIndex:
    """
        TermIndex is essentially keeps track of all index terms, their character
        offsets (ascending) and their word offsets in the document
    """

    def __init__(self, doc):
        self.document = doc
        self._index_terms = {}

        self._parse_document()

    def term_occurrences(self, term):
        """Return the occurrences of and indexed term, if found"""

        occurrences = None
        if term in self._index_terms:
            occurrences = self._index_terms[term]

        return occurrences

    # PRIVATE

    def _parse_document(self):
        """Parses the document to be indexed, identifying its Terms"""

        # Parse doc and build term index
        ignore = compile(DELIMITERS)
        index_term = None

        char_offset = 0
        word_offset = 0

        # Add space to delimit the document and simplify parsing
        doc_to_parse = self.document + ' '

        for c in doc_to_parse:
            if ignore.match(c) == None:
                if index_term == None:
                    # Start building our term
                    index_term = Term(c, char_offset, word_offset)
                    word_offset += 1
                else:
                    # Add to our term
                    index_term.term += c
            elif index_term:
                if index_term.term in self._index_terms:
                    # Subsequent Term appearances in the document
                    self._index_terms[index_term.term].append(index_term)
                else:
                    # Term appears for the first time
                    self._index_terms[index_term.term] = [index_term]

                index_term = None

            char_offset += 1

class Query:
    """
        Query represents a query against an index. In effect, this class represents
        a poorman's proximity search algorithm.
    """

    def __init__(self, query):
        self.original_query = query
        delimiters = compile(DELIMITERS)

        # Create a list of usable original query
        self.query_terms = filter(lambda x: len(x) > 0, delimiters.split(query.lower()))
        self.query_results = []

    def perform(self, term_index):
        """Performs the query on the provided term_index"""

        # No terms? No results
        if len(self.original_query) == 0:
            print "Nothing to search for."
            return

        # Try to find all terms in the index
        found_terms = []
        missing_term = False
        for query_term in self.query_terms:
            term_occurrences = term_index.term_occurrences(query_term)
            if term_occurrences:
                found_terms.extend(term_occurrences)
            else:
                missing_term = True
                break;

        # All terms must be found
        if missing_term:
            print "Not all terms were found."
            return

        best_match = self._best_match(found_terms)

        self._highlight(best_match, term_index.document)

    # PRIVATE
    def _best_match(self, found_terms):
        """Return best match phrase of all possible options"""

        combos = itertools.permutations(found_terms, len(self.query_terms))
        complete_results = filter(self._complete, combos)
        deduped_results = {}

        for result in complete_results:
            sorted_result = sorted(result, key=lambda x: x.word_offset)

            # Create a key for the deduped results dictionary
            key = ''
            for r in sorted_result:
                key += r.term + str(r.char_offset)

            if key not in deduped_results:
                self.query_results.append(sorted_result)
                deduped_results[key] = True

        best_distance = None
        best_match = None

        # The terms are now in order, we now calculate their total word
        # distance score by adding up each offset from the first to the
        # second word, second to third word, etc.
        #
        # lowest distance score wins. Ties are broken by the earliest match.

        for qr in self.query_results:
            total_distance = 0
            word_offset = qr[0].word_offset
            k = ''
            for idx, term in enumerate(qr):
                if idx > 0:
                    total_distance += term.word_offset - word_offset
                    word_offset = term.word_offset
                k += str(term) + ' '

            if best_distance == None or best_distance > total_distance:
                best_distance = total_distance
                best_match = qr # our best match so far

        return best_match

    def _complete(self, x):
        "Return True if the query terms consist of comlete unique terms"""

        keys = {}
        for t in x:
            keys[t.term] = t

        return len(keys) == len(self.query_terms)

    def _highlight(self, match, doc):
        """Simple highlighter indicating which block of text matched"""

        print ""
        print "Highlighted results below:"
        print ""

        start_offset = match[0].char_offset
        last_term = match[len(match) - 1]
        end_offset = last_term.char_offset + len(last_term.term)

        print doc[:start_offset] + '[' + \
            doc[start_offset:end_offset] + ']' + \
            doc[end_offset:]

def test():
    ti = TermIndex("""The George Washington Bridge in New York City is one of
 the oldest bridges ever constructed. It is now being remodeled because the
 bridge is a landmark. City officials say that the landmark bridge effort
 will create a lot of new jobs in the city.""" * 5 )

    default_query = "Landmark City Bridge"


    ti = TermIndex("a bbbbbbbbbb c d e f a b")
    default_query = "a b c"

    while True:
        print ""

        prompt = 'Enter your query (q to quit) [{}]: '.format(default_query)
        qt = "" # raw_input(prompt)

        if len(qt) == 0:
            qt = default_query
        else:
            qt = qt.strip()

        if qt == 'q':
            break;

        q = Query(qt)
        q.perform(ti)
        break


if __name__ == "__main__":
    test()
