#include <stdio.h>
#include <string.h>
#include <unistd.h>

typedef struct  {
    int progress;
    int min;
    char title[13];
} Status;

void timer(Status *status)
{
    printf("\033[H\033[J");
    for (int i = status->min; i >= 0 ; i--) {
        for (int j = 59; j >= 0 ; j--) {
            printf("%s\n", status->title);
            printf("%d:%d\n", i, j);
            sleep(1);
            printf("\033[H\033[J");
        }
    }
}

int main(void)
{
    Status status = {0};

    while (1) {
        switch (status.progress) {
            case 0: case 2: case 4: case 6:
                strcpy(status.title, "Focus");
                status.min = 24;
                break;
            case 1: case 3: case 5:
                strcpy(status.title, "Short Break!");
                status.min = 4;
                break;
            case 7:
                strcpy(status.title, "Long Break!");
                status.min = 14;
                break;
        }
        
        timer(&status);

        status.progress++;
        if (status.progress == 8) {
            status.progress = 0;
        }
    }

    return 0;
}
