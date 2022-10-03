# Vignere Cipher

The vigenere cipher is a classic caesar cipher algorithm that is used to encrypting and decrypting the text. The vigenere cipher is an algorithm of encrypting an alphabetic text that uses a series of interwoven caesar ciphers. It is based on a keyword's letters. It is an example of a polyalphabetic substitution cipher. This algorithm is easy to understand and implement. This algorithm was first described in 1553 by Giovan Battista Bellaso. It uses a Vigenere table or Vigenere square for encryption and decryption of the text. The vigenere table is also called the tabula recta.

# How Vignere Cipher Works

When the vigenere table is given, the encryption and decryption are done using the vigenere table (26 * 26 matrix) in this method.

![image](https://user-images.githubusercontent.com/74844505/193467139-32046e66-3be6-4188-80ff-19fd5fbba732.png)

**Example: The plaintext is "ILOVEYOU", and the key is "YERO".**

To generate a new key, the given key is repeated in a circular manner, as long as the length of the plain text does not equal to the new key.

## Encryption

![image](https://user-images.githubusercontent.com/74844505/193467435-af09902b-92cf-4e52-bfdc-0b15bd8dc6a1.png)

The first letter of the plaintext is combined with the first letter of the key. The column of plain text "J" and row of key "B" intersects the alphabet of "K" in the vigenere table, so the first letter of ciphertext is "K".

Similarly, the second letter of the plaintext is combined with the second letter of the key. The column of plain text "I" and row of key "Y" intersects the alphabet of "G" in the vigenere table, so the second letter of ciphertext is "P".

This process continues continuously until the plaintext is finished.

**Ciphertext = GPFJCCFI**

Encryption Formula:
$$E_i = (P_i + K_i)\mod26$$

## Decryption

Decryption is done by the row of keys in the vigenere table. First, select the row of the key letter, find the ciphertext letter's position in that row, and then select the column label of the corresponding ciphertext as the plaintext.

![image](https://user-images.githubusercontent.com/74844505/193467594-a61e8800-7f66-4002-94a6-ed206d06b4e2.png)

For example, in the row of the key is "Y and the ciphertext is "G" and this ciphertext letter appears in the column "I", that means the first plaintext letter is "I".

Next, in the row of the key is "E" and the ciphertext is "P" and this ciphertext letter appears in the column "L", that means the second plaintext letter is "L".

This process continues continuously until the ciphertext is finished.

**Plaintext = ILOVEYOU**

Decryption Formula:
$$D_i = (P_i - K_i + 26)\mod26$$

Source: https://www.javatpoint.com/vigenere-cipher
