import { Message } from "./message";

export interface ExecutionI {
  StartTime: Date;
  Messages: Message[];
}

export class Execution implements ExecutionI {
  StartTime: Date;
  Messages: Message[];

  constructor() {
    this.StartTime = new Date();
    this.Messages = [];
  }
}

export default class ExecutionManager {
  private ExecutionStore: Execution[];
  private CurrentExecution: Execution | null;

  constructor() {
    this.ExecutionStore = [];
    this.CurrentExecution = null;
  }

  NewExecution() {
    const newExecution = new Execution();
    this.ExecutionStore = [...this.ExecutionStore, newExecution];
    this.CurrentExecution = newExecution;
  }

  AddMessage(message: Message) {
    if (!this.CurrentExecution) {
      this.NewExecution();
    }
    this.CurrentExecution.Messages.push(message);
  }

  Clear() {
    this.ExecutionStore.length = 0; // This directly clears the array
    this.CurrentExecution = null;
  }

  GetExecutions(): Execution[] {
    return this.ExecutionStore;
  }
}
