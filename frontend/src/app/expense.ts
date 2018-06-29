export class Expense {
    constructor(
        public id: string,
        public amount: number,
        public accountID: string,
        public currency: string = 'EUR') { }
}
