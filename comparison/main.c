#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "misc.h"
#include "semiglobal.h"
#include "local.h"

void semiglobalExample(char *s1, char *s2){
	ll *result1 = NULL, *result2 = NULL;

	int s1Len = strlen(s1), s2Len = strlen(s2);

	int **matrix = allocEverything(s1Len, s2Len);

	fillSemiglobalMatrix(matrix, s1Len, s2Len, s1, s2);

	// printMatrix(matrix, s1Len, s2Len);
	// printf("\n");

	findSemiglobal(matrix, s1Len, s2Len, s1, s2, &result1, &result2);

	if(s1Len > s2Len){
		printf("%s\n", s1);
		printList(result2);
	} else{
		printf("%s\n", s2);
		printList(result1);
	}

	printf("\n");

	freeEverythingSemiglobal(matrix, s1Len, result1, result2);
}

void localExample(char *s1, char *s2){

	ll *results = NULL;

	int s1Len = strlen(s1), s2Len = strlen(s2);

	int **matrix = allocEverything(s1Len, s2Len);

	fillLocalMatrix(matrix, s1Len, s2Len, s1, s2);

	findLocals(matrix, s1Len, s2Len, &results);

	ll *i = results;

	while(i != NULL){
		printBacktrack(matrix, i->x, i->y, s1, s2);
		i = i->next;
	}

	freeEverythingLocal(matrix, s1Len, results);
}

int main(){
	char s1[] = "GTGCTCGGCGCACTTACCTCGGTGGGCCTGGTCCGTGCGATCCAGCCGTCGGGATCGGTCGCCCGC";
	char s2[] = "GTGGTCCGCGACTTTACCTCGGTGGGCCTGGC";

	char s3[] = "AATCCCGCGTCGGCGACAACCACCACCATGTCGTCTGCCGATCCTGCGGAGCGATTGCCGACATCGATTG";
	char s4[] = "GTCTGCCGATCCTGCGGAGCGATTGCCGACATCGATTGTTGCCCCAGCCGTC";

	char l1[] = "GTGCTCGGCGCACTTACCTCGGTGGGCCTGGTCCGTGCGATCCAGCCGTCGGGATCGGTCGCCCGC";
	char l2[] = "ATTTGCATGGTGGGCCTGAAAAGCAATGCCAGCCGTATGCGATCGAAATCTATGCCCGTCGGTGGG";

	char l3[] = "AATCCCGCGTCGGCGACAACCACCACCATGTCGTCTGCCGCGATCCAGCCGTCGGGATCGGTCGCCC";
	char l4[] = "TTCGATCGAACCACCACCAAGTCGTCTTAGCTTGAACTTCCAGCCGATTTCTCGCCC";

	printf("Semiglobal 1:\n");
	semiglobalExample(s1, s2);

	printf("Semiglobal 2:\n");
	semiglobalExample(s3, s4);

	printf("Local 1:\n");
	localExample(l1, l2);

	printf("Local 2:\n");
	localExample(l3, l4);
	return 0;
}
