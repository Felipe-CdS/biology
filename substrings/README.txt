Para rodar o programa utilize:

> go run .

ou compile e rode o executável com:

>go build .
>./substrings

#########################################################

Algumas informações sobre o programa:

A parte de subsequências calcula TODAS as subsequências possiveis como pedido no exercício.
Fazer isso cria alguns problemas computacionais e de tempo os quais tentei contornar com certas
mudanças.

A primeira mudança foi na entrada do problema. Para as substrings a entrada é feita com
o arquivo entry.txt que contém exatamente a fita de DNA informada no problema inicial. Já para
as subsequências, calcular tudo para uma fita daquele tamanho é totalmente inviável pela 
quantidade de possibilidades. Por esse motivo diminui a entrada para uma parcela da fita inicial,
dessa vez com apenas 70 caracteres. Essa nova entrada está no arquivo entry2.txt.

A segunda manobra para contornar os problemas é opicional, mas recomendada. Ao rodar o programa
é possivel passar algumas flags como argumento de entrada. As flags são -subseq, -all e -debug.

-subseq:
Essa flag limita em um número fixo a quantidade de subsequências encontradas. É possivel escolher
esse numero após a flag, por exemplo, escrevendo -subseq=10000 para as 100000 primeias subsequências.
O valor padrão do programa sem usar a flag é de 1000 subsequências.

-all:
Essa flag permite que todas as subsequências sejam encontradas sem limite, por isso use com cuidado.
Para a entrada do arquivo entry2.txt, por exemplo, a subsequência ATCCAGAATTCTCGGA tem 15394783 posições 
distintas e isso pode levar um tempo para ser calculado (~1min). Além disso, todas as posições encontradas 
são escritas no arquivo de output. Fique ciente de que esse arquivo pode tomar tamanhos consideravelmente 
grandes por isso.
O uso da flag é feito de forma boolena, logo usar -all é verdadeiro e omiti-la é falso.
obs.: Usar essa flag anula o uso da -subseq.

-debug:
Essa é uma flag opicional que escreve no output a base encontrada na fita ao invés da sua posição 
numérica. Serve apenas para verificar que as posições encontradas são realmente a subsequência,
pois só cria um arquivo com a mesma base repetida n vezes.

