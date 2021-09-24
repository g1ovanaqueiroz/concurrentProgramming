int buffer;
long counter = 0; //protocols starts with empty buffer

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
        put(i);
    }
}

void *consume(void *args)
{
    int i;
    while (1)
    {
        int tmp = get();
        prinf("%d\n", tmp);
    }
}