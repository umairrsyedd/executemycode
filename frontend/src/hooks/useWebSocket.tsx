import { Message, MessageType } from "@/types/message";
import { useState, useEffect, useRef } from "react";
import useWebSocket, { ReadyState, Options } from "react-use-websocket";

export const useCustomWebSocket = (url, onOutput, onDone, onError) => {
  const onOpen = (event) => {
    console.log("Connection Established With Server");
  };

  const onMessage = (event) => {
    const message: Message = JSON.parse(event.data);

    switch (message.type) {
      case MessageType.Output:
        onOutput(message.message);
        break;
      case MessageType.Done:
        onDone(message.message);
        break;
      case MessageType.Error:
        onError(message.message);
        break;
      default:
        console.log("Invalid Message Type from Server");
    }
  };

  const onClose = (event) => {
    console.log("Connection Closed With Server");
  };

  const webSocketOptions: Options = {
    onOpen: onOpen,
    onMessage: onMessage,
    onClose: onClose,
    reconnectAttempts: 0,
    reconnectInterval: 1000,
    retryOnError: true,
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
