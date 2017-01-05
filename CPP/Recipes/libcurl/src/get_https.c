/*************************************************************************
 > File Name: get_https.c
 > Author:
 > Mail:
 > Created Time: 2016年12月29日 星期四 11时33分51秒
 ************************************************************************/

#include <curl/curl.h>
#include <stdio.h>
#include <string.h>

#ifndef HTTPS_GET_URL
#define HTTPS_GET_URL "https://192.168.40.161/api/v1/maintenance/lineinfo"
#endif

size_t
writeFunction(void *ptr, size_t size, size_t nmemb, char* data)
{
    memcpy(data, (char *)ptr, size*nmemb);

    return size * nmemb;
}

int
main(int argc, char **argv)
{
    CURL *curl;
    CURLcode res;

    curl_global_init(CURL_GLOBAL_DEFAULT);

    curl = curl_easy_init();
    if (curl) {
        curl_easy_setopt(curl, CURLOPT_URL, HTTPS_GET_URL);
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, 0L);
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, 0L);

        char response_string[1024];
        char header_string[1024];
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writeFunction);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, response_string);
        // curl_easy_setopt(curl, CURLOPT_HEADERDATA, header_string);
        curl_easy_setopt(curl, CURLOPT_VERBOSE, 1);

        res = curl_easy_perform(curl);
        /* Check for errors */
        if(res != CURLE_OK)
            fprintf(stderr, "curl_easy_perform() failed: %s\n",
                    curl_easy_strerror(res));

        printf("recv: %s\n", response_string);
        /* always cleanup */
        curl_easy_cleanup(curl);
    }

    curl_global_cleanup();

    return 0;
}
