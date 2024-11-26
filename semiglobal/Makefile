CC=gcc
NAME=semiglobal
CFLAGS=-Wall -Wextra -Werror -I.-
SRCS=main.c misc.c
OBJS = $(SRC:%.c=%.o)

all: $(NAME)

clean:

fclean: clean
	rm $(NAME)

re: fclean all

$(NAME): $(SRCS) 
	$(CC) $(SRCS) -o $(NAME) $(CFLAGS)

.PHONY: all clean fclean re

