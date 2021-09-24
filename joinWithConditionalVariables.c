// this is pseudo C code, so comments are needed
// int done = 0;
// mutex_t mutex;
// cond_t cond; //declaring condicional var
// void run(void *arg) {
//     lock(mutex);
//     printf("hey mom\n");
//     signal(cond);
//     unlock(mutex);
// }

// int main()
// {
//     pthread_t child;
//     pthread_create(&child, NULL, run, NULL); //create thread
//     lock(mutex);
//     while(done == 0) {
//         wait(cond);
//     }
//     unlock(mutex);
//     return 0;
// }