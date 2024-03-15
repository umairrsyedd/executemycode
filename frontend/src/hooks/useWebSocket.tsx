import { Message, MessageType } from "@/types/message";
import { useState, useEffect, useRef } from "react";
import useWebSocket, { ReadyState, Options } from "react-use-websocket";

export const useCustomWebSocket = (
  url,
  onOutput,
  onDone,
  onError,
  onConnected,
  onReconnectStop
) => {
  const onOpen = (event) => {
    console.log("Connection Established With Server");
  };

  const onMessage = (event) => {
    const message: Message = JSON.parse(event.data);

    switch (message.type) {
      case MessageType.Output:
        onOutput(message);
        break;
      case MessageType.Done:
        onDone(message);
        break;
      case MessageType.Error:
        onError(message);
        break;
      default:
        console.log("Invalid Message Type from Server");
    }
  };

  const onClose = (event) => {
    console.log("Connection Closed With Server");
  };

  const onSocketError = (event) => {
    console.log("An Error Occured Bro");
  };

  const webSocketOptions: Options = {
    onOpen: onConnected,
    onMessage: onMessage,
    onClose: onClose,
    onError: onSocketError,
    reconnectAttempts: 20,
    reconnectInterval: 2000,
    retryOnError: true,
    onReconnectStop: onReconnectStop,
  };

  const { sendJsonMessage, lastMessage, readyState, sendMessage } =
    useWebSocket(
      process.env.NEXT_PUBLIC_EXECUTION_SERVER_URL,
      webSocketOptions
    );

  const sendCode = (code: string, language: string) => {
    const message: Message = {
      type: MessageType.Code,
      message: code,
      language: language,
    };
    sendJsonMessage(message);
  };

  const sendInput = (input: string) => {
    const message: Message = {
      type: MessageType.Input,
      message: input,
    };
    sendJsonMessage(message);
  };

  const sendClose = () => {
    const message: Message = {
      type: MessageType.Close,
    };
    sendJsonMessage(message);
  };

  return {
    onMessage,
    sendCode,
    sendInput,
    sendClose,
  };
};

export default useWebSocket;
