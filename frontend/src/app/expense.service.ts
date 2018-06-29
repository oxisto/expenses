import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Expense } from './expense';

const EXPENSE_ENDPOINT = '/api/expenses';

@Injectable({
  providedIn: 'root'
})
export class ExpenseService {

  constructor(private http: HttpClient) { }

  getExpenses(): Observable<Expense[]> {
    return this.http.get<Expense[]>(EXPENSE_ENDPOINT);
  }

  getExpense(id): Observable<Expense> {
    return this.http.get<Expense>(EXPENSE_ENDPOINT + '/' + id);
  }

  postExpense(expense: Expense): Observable<Expense> {
    return this.http.post<Expense>(EXPENSE_ENDPOINT, expense);
  }

  putExpense(expense: Expense): Observable<Expense> {
    return this.http.put<Expense>(EXPENSE_ENDPOINT + '/' + expense.id, expense);
  }

}
