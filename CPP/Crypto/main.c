/*************************************************************************
 > File Name: main.c
 > Author:
 > Mail:
 > Created Time: 2016年12月15日 星期四 16时23分49秒
 ************************************************************************/
#include <stdio.h>

#include <sodium.h>

#define MESSAGE (const unsigned char *) "test"
#define MESSAGE_LEN 4

#define ADDITIONAL_DATA (const unsigned char *) "123456"
#define ADDITIONAL_DATA_LEN 6

#define SODIUM_NONCE_LEN   8U
#define SODIUM_KEY_LEN    32U
#define SODIUM_CIPHER_LEN 16U

int
main(void)
{
    if (sodium_init() != 0) {
        return 1;
    }

    printf("init success!\n");

    unsigned char nonce[SODIUM_NONCE_LEN];
    unsigned char key[SODIUM_KEY_LEN];
    unsigned char ciphertext[MESSAGE_LEN + SODIUM_CIPHER_LEN];
    unsigned long long ciphertext_len;

    randombytes_buf(key, sizeof(key));
    randombytes_buf(nonce, sizeof(nonce));

    crypto_aead_chacha20poly1305_encrypt(ciphertext, &ciphertext_len, MESSAGE, MESSAGE_LEN, \
            ADDITIONAL_DATA, ADDITIONAL_DATA_LEN, NULL, nonce, key);
    int index;
    for (index = 0; index < ciphertext_len; index++) {
        printf("0x%02X ", ciphertext[index]);
    }
    printf("\n");

    unsigned char decrypted[MESSAGE_LEN];
    unsigned long long decrypted_len;

    if (crypto_aead_chacha20poly1305_decrypt(decrypted, &decrypted_len, NULL, \
            ciphertext, ciphertext_len, ADDITIONAL_DATA, ADDITIONAL_DATA_LEN, \
            nonce, key) != 0) {
        printf("decrypt fail!\n");
        return 1;
    }
    printf("message: %s\n", decrypted);

    return 0;
}
