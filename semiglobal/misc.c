#include "misc.h"
#include <stdio.h>
#include <stdlib.h>

// simple linked list
ll* addNewNode(ll *l, char v){
	ll *new = (ll*) malloc(sizeof(ll));
	new->value = v;
	new->next = NULL;

	if(l == NULL) return new;

	new->next = l;
	return new;
}

void freeList(ll *l){
	if(l && l->next) freeList(l->next);
	free(l);
}

void printList(ll *l){
	if(l){
		printf("%c", l->value);
		printList(l->next);
	}
}

int** allocEverything(int s1Len, int s2Len){

	int **holder = (int**) malloc((s1Len + 1) * sizeof(int*));

	for(int i = 0; i <= s1Len; i++){
		int *e = (int*) malloc((s2Len + 1) * sizeof(int));

		for(int j = 0; j <= s2Len; j++)
			*(e+j) = 0;

		*(holder+i) = e;
	}

	return holder;
}

void freeEverything(int** m, int s1Len, ll* r1, ll* r2){
	for(int i = 0; i <= s1Len; i++) 
		free(m[i]);
	free(m);

	freeList(r1);
	freeList(r2);
}

void printMatrix(int** m, int s1Len, int s2Len){
	for(int i = 0; i <= s1Len; i++){
		for(int j = 0; j <= s2Len; j++){
			printf("%2d ", m[i][j]);
		}
		printf("\n");
	}
}

int max(int x, int y, int z){
	int result = x;
	if(y > result) result = y;
	if(z > result) result = z;
	return result;
}
