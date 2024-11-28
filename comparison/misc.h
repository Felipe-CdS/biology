#ifndef MISC_H_INCLUDED 
#define MISC_H_INCLUDED

struct lls {
	char value;
	int x, y;
	struct lls *next;
};

typedef struct lls ll;

ll* addNewSemiglobalNode(ll *l, char v);
ll* addNewLocalNode(ll *l, int x, int y);

void printMatrix(int** m, int s1Len, int s2Len);
void printList(ll *l);
void printLocalsList(ll *l);

int** allocEverything(int s1Len, int s2Len);

void freeList(ll *l);
void freeEverythingSemiglobal(int** m, int s1Len, ll* r1, ll* r2);
void freeEverythingLocal(int** m, int s1Len, ll* r);

int max(int x, int y, int z);

#endif

