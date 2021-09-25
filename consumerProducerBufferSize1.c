#include <pthread.h>

//this code still has a deadlock problem

int buffer;
long counter = 0; //protocols starts with empty buffer
pthread_cond_t cond_full;
pthread_cond_t cond_empty;
pthread_mutex_t mutex;

void put(int value)
{
    assert(counter == 0);
    counter = 1;
    buffer = value;
}

int get()
{
    assert(counter == 1);
    counter = 0;
    return buffer;
}

void *produce(void *args)
{
    int i;
    int items_to_produce = (int)args;
    for (i = 0; i < items_to_produce; i++)
    {
        lock(mutex);
        while (counter == 1)
        {
            wait(cond_full, mutex);
        }
        put(i);
        signal(cond_empty);
        unlock(mutex);
    }
    return NULL;
}

void *consume(void *args)
{
    int i;
    while (1)
    {
        lock(mutex);
        while (counter == 0)
        {
            wait(cond_empty, mutex);
        }
        int tmp = get();
        signal(cond_full);
        unlock(mutex);
    }
    return NULL;
}