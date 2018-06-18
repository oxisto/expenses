export class Expense {
    constructor(
        public amount: number,
        public accountID: number,
        public currency: string = 'EUR') { }
}
