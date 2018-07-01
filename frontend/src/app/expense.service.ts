import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { AuthService } from './auth.service';
import { Expense } from './expense';

const EXPENSE_ENDPOINT = '/api/expenses';

@Injectable({
  providedIn: 'root'
})
export class ExpenseService {

  constructor(private http: HttpClient, private authService: AuthService) { }

  getExpenses(): Observable<Expense[]> {
    return this.http.get<Expense[]>(EXPENSE_ENDPOINT)
      .pipe(
        catchError(this.authService.handleHttpError.bind(this.authService))
      );
  }

  getExpense(id): Observable<Expense> {
    return this.http.get<Expense>(EXPENSE_ENDPOINT + '/' + id)
      .pipe(
        catchError(this.authService.handleHttpError.bind(this.authService))
      );
  }

  postExpense(expense: Expense): Observable<Expense> {
    return this.http.post<Expense>(EXPENSE_ENDPOINT, expense)
      .pipe(
        catchError(this.authService.handleHttpError.bind(this.authService))
      );
  }

  putExpense(expense: Expense): Observable<Expense> {
    return this.http.put<Expense>(EXPENSE_ENDPOINT + '/' + expense.id, expense)
      .pipe(
        catchError(this.authService.handleHttpError.bind(this.authService))
      );
  }

}
