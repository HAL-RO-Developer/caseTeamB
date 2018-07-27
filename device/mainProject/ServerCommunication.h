#ifndef __SERVER_COMMUNICATION_H__
#define __SERVER_COMMUNICATION_H__

/* --- includeファイル --- */
#include "ArduinoLibrary.h"  
#include "System.h"
#include "InfoStruct.h"

#define SERVER_SUCCESS      (   1 )
#define SERVER_CONNECT_ERROR  (  -1 )

/* --- HTTPステータスコード --- */
#define HTTP_STATUS_OK        ( 200 )
#define HTTP_STATUS_FORBIDDEN ( 403 )
#define HTTP_STATUS_NOT_FOUND ( 404 )

class ServerCommunication{
  private:
    WiFiClientSecure client;
    String host;
      String port;
    String url;
    SINT status;
  public:
    ServerCommunication(String host="");
    SSHT connect(String host, String port);
    void post(String url, String data, String host);
    void put(String url, String data, String host);
    void get(String url, String data);
    SINT response_serv(SINT* data);
    SINT response_dev(String data[]);
    void setUrl(String url);
    SINT getStatus();
};

#endif /* __SERVER_COMMUNICATION_H__ */
