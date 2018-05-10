#include <stdio.h>
#include <string.h>

const int MAXLEN = 100000;

int ss(char *s) {
    char *p, *a, *b;
    int n = strlen(s);
    if (s[n-1] == '\n') {
        s[--n] = '\0';
    }
    int c = n;
    for (p=s+1; *p; ++p) {
        for (a=s, b=p; *a == *b; ++a, ++b) {
            ++c;
        }
    }
    return c;
}

int main() {
    char buf[MAXLEN+1];
    if (!fgets(buf, sizeof(buf)-1, stdin)) {
        return 0;
    }
    while (fgets(buf, sizeof(buf)-1, stdin)) {
        printf("%d\n", ss(buf));
    }
    return 0;
}
