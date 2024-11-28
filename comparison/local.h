#ifndef LOCAL_H_INCLUDED 
#define LOCAL_H_INCLUDED

void fillLocalMatrix(int** m, int s1len, int s2len, char* s1, char* s2);

void findLocals(int** m, int c, int r, ll** results);

void printBacktrack(int** m, int x, int y, char* s1, char* s2);
#endif
