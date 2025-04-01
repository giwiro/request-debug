export class NotifiableError extends Error {
  constructor(msg: string) {
    super(msg);

    Object.setPrototypeOf(this, NotifiableError.prototype);
  }
}
