#include "misc.h"

void fillSemiglobalMatrix(int** m, int s1len, int s2len, char* s1, char* s2){

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

			m[i][j] = max(t, d, l);

			++j;
		}
		j = 1;
		++i;
	}
}

void findSemiglobal(int** m, int c, int r, char* s1, char* s2, ll** r1, ll** r2){

	int i = 0, j = 0;
	int max = m[c][r], maxX = c, maxY = r;

	// find max sim on last col
	for(i = 0; i <= c; i++){
		if(m[i][r] > max){
			max = m[i][r];
			maxX = i;
			maxY = r;
		}
	}

	// find max sim on last row
	for(i = 0; i <= r; i++){
		if(m[c][i] > max){
			max = m[c][i];
			maxX = c;
			maxY = i;
		}
	}

	//add end spaces
	for(i = 0; i < (c - maxX - 1); i++){
		if(c < r){
			*r1 = addNewSemiglobalNode(*r1, '-');
		} else {
			*r2 = addNewSemiglobalNode(*r2, '-');
		}
	}

	i = maxX;
	j = maxY;

	while(1){
		int t = m[i-1][j];
		int d = m[i-1][j-1];
		// int l = m[i][j-1];


		if(m[i][j] == (t-2)){
			*r1 = addNewSemiglobalNode(*r1, '-');
			*r2 = addNewSemiglobalNode(*r2, s2[j-1]);
			--j;
		} else if(m[i][j] == (d-1) || m[i][j] == (d+1)) {
			*r1 = addNewSemiglobalNode(*r1, s1[i-1]);
			*r2 = addNewSemiglobalNode(*r2, s2[j-1]);
			--j;
			--i;
		} else {
			*r1 = addNewSemiglobalNode(*r1, s1[i-1]);
			*r2 = addNewSemiglobalNode(*r2, '-');
			--i;
		}

		if(i == 0){
			while (j > 0) {
				*r1 = addNewSemiglobalNode(*r1, '-');
				*r2 = addNewSemiglobalNode(*r2, s2[j-1]);
				--j;
			}
			break;
		}

		if(j == 0){
			while (i > 0){
				*r2 = addNewSemiglobalNode(*r2, '-');
				*r1 = addNewSemiglobalNode(*r1, s1[i-1]);
				--i;
			}
			break;
		}
	}
}
