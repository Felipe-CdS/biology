#ifndef SEMIGLOBAL_H_INCLUDED 
#define SEMIGLOBAL_H_INCLUDED

void fillSemiglobalMatrix(int** m, int s1len, int s2len, char* s1, char* s2);
void findSemiglobal(int** m, int c, int r, char* s1, char* s2, ll** r1, ll** r2);

#endif
