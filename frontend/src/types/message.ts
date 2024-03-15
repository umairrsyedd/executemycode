export enum MessageType {
  // Sent From Client
  Code = "code",
  Input = "input",
  Close = "close",

  // Sent From Server
  Output = "output",
  Done = "done",
  Error = "error",
}

export interface Message {
type : MessageType,
message: string,
language: string,
}