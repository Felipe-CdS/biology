#ifndef MISC_H_INCLUDED 
#define MISC_H_INCLUDED

struct lls {
	char value;
	struct lls *next;
};

typedef struct lls ll;

ll* addNewNode(ll *l, char v);
void freeList(ll *l);
void printList(ll *l);

int** allocEverything(int s1Len, int s2Len);
void freeEverything(int** m, int s1Len, ll* r1, ll* r2);
void printMatrix(int** m, int s1Len, int s2Len);
int max(int x, int y, int z);

#endif

