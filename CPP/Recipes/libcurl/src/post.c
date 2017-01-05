/***************************************************************************
 *                                  _   _ ____  _
 *  Project                     ___| | | |  _ \| |
 *                             / __| | | | |_) | |
 *                            | (__| |_| |  _ <| |___
 *                             \___|\___/|_| \_\_____|
 *
 * Copyright (C) 1998 - 2015, Daniel Stenberg, <daniel@haxx.se>, et al.
 *
 * This software is licensed as described in the file COPYING, which
 * you should have received as part of this distribution. The terms
 * are also available at https://curl.haxx.se/docs/copyright.html.
 *
 * You may opt to use, copy, modify, merge, publish, distribute and/or sell
 * copies of the Software, and permit persons to whom the Software is
 * furnished to do so, under the terms of the COPYING file.
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTY OF ANY
 * KIND, either express or implied.
 *
 ***************************************************************************/
/* <DESC>
 * simple HTTP POST using the easy interface
 * </DESC>
 */
#include <stdio.h>
#include <curl/curl.h>
const char *json = "{\"agent_ip\":\"127.0.0.1\",\"linecount\":3, \
    \"lineinfo\":[{ \
        \"vps_uuid\":\"04b84af36ece4d6b58d3a13a9f26fe60\", \
        \"ping\":72 \
    }, { \
    \"vps_uuid\":\"1aebbe2629d3f3673c559c281510d758\", \
    \"ping\":72 \
    }, { \
    \"vps_uuid\":\"3802d31f31e7d95c4229eb47d875f77c\", \
    \"ping\":0 \
    }] \
}";
int main(void)
{
  CURL *curl;
  CURLcode res;

  /* In windows, this will init the winsock stuff */
  curl_global_init(CURL_GLOBAL_ALL);

  /* get a curl handle */
  curl = curl_easy_init();
  if(curl) {
    /* First set the URL that is about to receive our POST. This URL can
       just as well be a https:// URL if that is what should receive the
       data. */
    curl_easy_setopt(curl, CURLOPT_URL, "http://192.168.40.161:5000/api/v1/maintenance/ping");
     struct curl_slist *headers = NULL;
    // curl_slist_append(headers, "Expect:");
    // curl_slist_append(headers, "Accept: application/json");
    curl_slist_append(headers, "Content-Type: application/json");
    // POST
    // curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, "POST");
    // set header
    curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
    /* Now specify the POST data */
    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, json);

    /* Perform the request, res will get the return code */
    res = curl_easy_perform(curl);
    /* Check for errors */
    if(res != CURLE_OK)
      fprintf(stderr, "curl_easy_perform() failed: %s\n",
              curl_easy_strerror(res));
    else
        printf("CURLE_OK\n");

    /* always cleanup */
    curl_easy_cleanup(curl);
  }
  curl_global_cleanup();
  return 0;
}


