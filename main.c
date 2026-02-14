// todo: pause button
// todo: restart button
// todo: sound

#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <libnotify/notify.h>
#include <ncurses.h>

typedef struct  {
    int progress;
    int min;
    char title[20];
    char notify_title[13];
    char message[30];
} Status;

void table(int y1, int x1, int y2, int x2)
{
    mvhline(y1, x1, 0, x2-x1);
    mvhline(y2, x1, 0, x2-x1);
    mvvline(y1, x1, 0, y2-y1);
    mvvline(y1, x2, 0, y2-y1);
    mvhline(2, x1, 0, x2-x1);
    mvvline(1, 14, 0, 1);

    mvaddch(y1, x1, ACS_ULCORNER);
    mvaddch(y2, x1, ACS_LLCORNER);
    mvaddch(y1, x2, ACS_URCORNER);
    mvaddch(y2, x2, ACS_LRCORNER);
    mvaddch(2, x1, ACS_LTEE);
    mvaddch(2, x2, ACS_RTEE);
}

void timer(Status *status)
{
    for (int i = status->min; i >= 0 ; i--) {
        for (int j = 59; j >= 0 ; j--) {
            mvprintw(1, 1, "%s\n", status->title);
            mvprintw(3, 5, "%02d:%02d\n", i, j);

            table(0, 0, 4, 14);

            refresh();
            sleep(1);
        }
    }
}

int main(void)
{
    initscr();
    curs_set(0); 

    Status status = {0};

    while (1) {
        switch (status.progress) {
            case 0: case 2: case 4: case 6:
                strcpy(status.title, "    Focus    ");
                strcpy(status.notify_title, "Focus");
                strcpy(status.message, "Time to focus!");
                status.min = 24;
                break;
            case 1: case 3: case 5:
                strcpy(status.title, " Short Break ");
                strcpy(status.notify_title, "Short Break");
                strcpy(status.message, "Time for a short break!");
                status.min = 4;
                break;
            case 7:
                strcpy(status.title, " Long Break  ");
                strcpy(status.notify_title, "Long Break");
                strcpy(status.message, "Time for a long break!");
                status.min = 14;
                break;
        }

        /* Show notification */
        NotifyNotification *notif;
        notify_init("Pomodoro");
        notif = notify_notification_new(status.notify_title, status.message, NULL);
        notify_notification_show(notif, NULL);

        timer(&status);

        status.progress++;
        if (status.progress == 8) {
            status.progress = 0;
        }
    }

	endwin();	

    return 0;
}
