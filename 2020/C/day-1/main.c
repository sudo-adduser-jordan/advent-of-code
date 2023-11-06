#include <stdio.h>
#include <stdlib.h>
#include "../lib/color.c"
#include "../lib/array.c"

const char file[] = "input.txt";

char line[10];
int lineLen = 10;
int linePointer;

int results[250];
int resultsLen = 250;
int resultPointer;

int total;
int number;
int result;

int PartOne()
{
    FILE *filePointer;
    filePointer = fopen(file, "r");
    if (filePointer != NULL)
    {
        while (fgets(line, 100, filePointer))
        {
            if (line[0] == '\n')
            {

                results[resultPointer] = total;
                resultPointer++;
                total = 0;
            };

            number = atoi(line);
            total += number;
        }
    }
    else
    {
        printf("Not able to open the file.");
    }

    result = FindGreatestValue(results, resultsLen);

    PrintArray(results, resultsLen);
    printf(MAGENTA_LABEL "result: " RESET_COLOR);
    printf(GREEN_TEXT "%d\n" RESET_COLOR, result);

    fclose(filePointer);
    return 0;
}

int PartTwo()
{
    FILE *filePointer;
    filePointer = fopen(file, "r");
    if (filePointer != NULL)
    {
        while (fgets(line, 100, filePointer))
        {
            if (line[0] == '\n')
            {

                results[resultPointer] = total;
                resultPointer++;
                total = 0;
            };

            number = atoi(line);
            total += number;
        }
    }
    else
    {
        printf("Not able to open the file.");
    }

    qsort(results, resultsLen, sizeof(int), CompareFunction);
    result = results[247] + results[248] + results[249];

    PrintArray(results, resultsLen);
    printf(MAGENTA_LABEL "result: " RESET_COLOR);
    printf(GREEN_TEXT "%d" RESET_COLOR, result);

    fclose(filePointer);
    return 0;
}

int main()
{
    // PartOne();
    PartTwo();

    return 0;
}
