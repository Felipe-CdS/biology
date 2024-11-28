#include <stdio.h>
#include <string.h>
#include "misc.h"

void fillLocalMatrix(int** m, int s1len, int s2len, char* s1, char* s2){

	int i = 1, j = 1;

	while(i <= s1len){
		while(j <= s2len){

			int t = m[i-1][j];
			int d = m[i-1][j-1];
			int l = m[i][j-1];


			if(s1[i-1] == s2[j-1]) ++d;
			else --d;

			t -= 2;
			l -= 2;

			if(max(t, d, l) > 0){
				m[i][j] = max(t, d, l);
			} else {
				m[i][j] = 0;
			}

			++j;
		}
		j = 1;
		++i;
	}
}

void findLocals(int** m, int c, int r, ll** results){

	int highest = 0;

	// find max sim 
	for(int i = 0; i <= c; i++){
		for(int j = 0; j <= r; j++){
			if(m[i][j] > highest)
				highest = m[i][j];
		}
	}

	// printf("sim: %d\n", highest);

	for(int i = 0; i <= c; i++){
		for(int j = 0; j <= r; j++){
			if(m[i][j] == highest)
				*results = addNewLocalNode(*results, j, i);
		}
	}

}

void printBacktrack(int** m, int x, int y, char* s1, char* s2){

	int s1Len = strlen(s1), s2Len = strlen(s2);
	char r1[s1Len], r2[s2Len];
	int idx = 0;

	while(1){
		if(m[y][x] == 0) break;

		int t = m[y-1][x];
		int d = m[y-1][x-1];
		int l = m[y][x-1];

		if(m[y][x] == (t-2)){
			r1[idx] = '-';
			r2[idx] = s2[y];
			idx++;
			--x;
		} else if(m[y][x] == (d-1) || m[y][x] == (d+1)) {
			r1[idx] = s1[x];
			r2[idx] = s2[y];
			idx++;
			--x;
			--y;
		} else if(m[y][x] == (l-2)){
			r1[idx] = s1[x];
			r2[idx] = '-';
			idx++;
			--y;
		} else {
			break;
		}
	}

	for(int i = (idx-1); i >= 0; i--){
		printf("%c", r1[i]);
	}

	printf("\n");

	for(int i = (idx-1); i >= 0; i--){
		printf("%c", r2[i]);
	}

	printf("\n");
}
