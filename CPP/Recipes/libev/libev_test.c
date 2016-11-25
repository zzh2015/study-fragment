/*************************************************************************
  > File Name: libev_test.c
  > Author:
  > Mail:
  > Created Time: 2016年11月21日 星期一 16时44分59秒
 ************************************************************************/
#include <ev.h>
#include <stdio.h>


static void stdin_cb (EV_P_ ev_io *w, int revents) {
    puts("stdin ready");

    ev_io_stop(EV_A_ w);

    ev_break(EV_A_ EVBREAK_ALL);
}

static void timeout_cb (EV_P_ ev_timer *w, int revents) {
    puts("timeout");

    ev_break (EV_A_ EVBREAK_ONE);
}

int main() {
    printf ("major:%d minor:%d, time: %f\n", ev_version_major(), ev_version_minor(), \
            ev_time());
    ev_sleep(1.0);
    ev_io stdin_watcher;
    ev_timer timeout_watcher;
    struct ev_loop *loop = EV_DEFAULT;


    ev_io_init (&stdin_watcher, stdin_cb, 1, EV_READ);
    ev_io_start (loop, &stdin_watcher);

    ev_timer_init(&timeout_watcher, timeout_cb, 5.5, 0.);
    ev_timer_start(loop, &timeout_watcher);

    ev_run(loop, 0);

    return 0;
}
