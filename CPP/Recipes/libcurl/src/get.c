#include <curl/curl.h>
#include <stdio.h>
#include <string.h>

size_t writeFunction(void *ptr, size_t size, size_t nmemb, char* data) {
    // data->append((char*) ptr, size * nmemb);
    memcpy(data, (char *)ptr, size*nmemb);

    return size * nmemb;
}

int main(int argc, char** argv) {
    CURL *curl = curl_easy_init();
    if (curl) {
        curl_easy_setopt(curl, CURLOPT_URL, "https://192.168.40.161:5000/api/v1/maintenance/lineinfo");
        // curl_easy_setopt(curl, CURLOPT_NOPROGRESS, 1L);
        // curl_easy_setopt(curl, CURLOPT_USERPWD, "user:pass");
        // curl_easy_setopt(curl, CURLOPT_USERAGENT, "curl/7.42.0");
        // curl_easy_setopt(curl, CURLOPT_MAXREDIRS, 50L);
        // curl_easy_setopt(curl, CURLOPT_TCP_KEEPALIVE, 1L);

        char response_string[1024];
        char header_string[1024];
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writeFunction);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, response_string);
        curl_easy_setopt(curl, CURLOPT_HEADERDATA, header_string);
        curl_easy_setopt(curl, CURLOPT_VERBOSE, 1);

        char* url;
        long response_code;
        double elapsed;
        curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &response_code);
        curl_easy_getinfo(curl, CURLINFO_TOTAL_TIME, &elapsed);
        curl_easy_getinfo(curl, CURLINFO_EFFECTIVE_URL, &url);

        curl_easy_perform(curl);
        curl_easy_cleanup(curl);
        curl = NULL;
    }
}
