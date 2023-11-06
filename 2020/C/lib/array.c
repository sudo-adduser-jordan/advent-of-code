#include <stdio.h>

int PrintArray(int array[], int len)
{
    int loop;
    for (loop = 0; loop < len; loop++)
    {
        printf("%d ", array[loop]);
    }
    printf("\n");
    return 0;
}

int CompareFunction(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
};

int FindGreatestValue(int array[], int len)
{
    int x;
    for (int i = 0; i < len; i = i + 1)
    {
        if (x < array[i])
        {
            x = array[i];
        };
    }
    return x;
};
