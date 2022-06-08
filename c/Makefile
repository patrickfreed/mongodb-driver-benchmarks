scratch: main.o
	gcc -o scratch main.o `pkg-config --cflags --libs libmongoc-1.0` -lpthread -ldl -lm

main.o: main.c
	gcc -c main.c `pkg-config --cflags --libs libmongoc-1.0`

clean:
	rm scratch main.o

test:
	./scratch
