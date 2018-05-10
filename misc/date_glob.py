import calendar
import datetime

class DateGlobber():
    '''
        given date_from=2013-08-21 and date_from=2013-09-02, generate:

        2013{0821,0822,0823,0824,0825,0826,0827,0828,0829,0830,0831,0901,0902}
    '''

    def date_range(self,start_date, end_date):
        '''
          generate series of dates in range from start_date to end_date inclusive
        '''
        for n in range(int ((end_date - start_date).days + 1)):
            yield start_date + datetime.timedelta(n)

    def date_glob(self,date_from,date_to):
        '''
           given date_from=2013-08-21 and date_from=2013-09-02, generate:

           2013{0821,0822,0823,0824,0825,0826,0827,0828,0829,0830,0831,0901,0902}
        '''
        if date_to==date_from: return date_from.strftime('%Y%m%d')
        if date_to < date_from: raise Error("Invalid date range date_from:"+date_from+" date_to"+date_to)
        same_year = (date_from.year==date_to.year)
        same_month = (date_from.month==date_to.month)
        d = ""
        if same_year and same_month:
           d += date_from.strftime('%Y%m%d')[:6]
        else:
           if same_year:
               d += str(date_from.year)
        a=[]
        for r in self.date_range(date_from,date_to):
            dt = r.strftime('%Y%m%d')
            if same_year and same_month:
                a.append( str(dt)[6:] )
            else:
                if same_year:
                    a.append( str(dt)[4:] )
                else:
                    a.append(dt)
        return d + "{" + ",".join(a) + "}"

def date_range(self,start_date, end_date):
    for n in range(int ((end_date - start_date).days + 1)):
        yield start_date + datetime.timedelta(n)


def date_glob(first, last):
    header = ''
    dayfmt = '{0.year}{0.month:02d}{0.day:02d}'

    if first.year == last.year:
        if first.month == last.month:
            header = '{0.year}{0.month:02d}'.format(first)
            dayfmt = '{0.day:02d}'
        else:
            header = '{0.year}'.format(first)
            dayfmt = '{0.month:02d}{0.day:02d}'

    days = (dayfmt.format(d) for d in date_range(first, last))
    return '%s{%s}' % (header, ','.join(days))


first = datetime.datetime(2013, 8, 21)
last = datetime.datetime(2013, 9, 3)

print month_glob(first, last)

g = DateGlobber()
print g.date_glob(first, last)
