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
    char title[13];
    char message[30];
} Status;

void timer(Status *status)
{
    for (int i = status->min; i >= 0 ; i--) {
        for (int j = 59; j >= 0 ; j--) {
            mvprintw(0, 0, "%s\n", status->title);
            mvprintw(1, 0, "%02d:%02d\n", i, j);
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
                strcpy(status.title, "Focus");
                strcpy(status.message, "Time to focus!");
                status.min = 24;
                break;
            case 1: case 3: case 5:
                strcpy(status.title, "Short Break!");
                strcpy(status.message, "Time for a short break!");
                status.min = 4;
                break;
            case 7:
                strcpy(status.title, "Long Break!");
                strcpy(status.message, "Time for a long break!");
                status.min = 14;
                break;
        }

        /* Show notification */
        NotifyNotification *notif;
        notify_init("Pomodoro");
        notif = notify_notification_new(status.title, status.message, NULL);
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
