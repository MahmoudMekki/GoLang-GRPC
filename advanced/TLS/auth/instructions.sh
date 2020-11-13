#Generating the CA and tls files for the client and server 

# o/p files 
# ca.key : certificate authority private key file (this shouldn`t be shared in real life)
# ca.crt : certificate authority rust certificate (this should be shared with the users in real life)
# server.key : server private key password protected (this shouldn`t be shared )
# server.csr :server certificate signing request (this should be shared )
# server.crt : server certificate signed by the CA (this wouldbe sent back by the CA owner)
# server.pem : conversion of server.key to server.pem that grpc likes (this shouldn`t be shared)


#summary
# private files : ca.key , server.key , server.pem , server.crt
#shared files : ca.crt(needed by the client) , server.csr (needede by the the CA)


#this could be change to any enviroment you want!!! 
# for example mahmoud.com

SERVER_CN=localhost

#step 1 : Generate the Certificate Authority + trust certificate
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"

#step2 : generate the server private key (server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

#setp3 : get a certificate signing from ca (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"

#step4 : sign the certificate with the CA we created (it's called self signing) server.crt
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

#step5 : convert ther server ccertificate to .pem format
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem