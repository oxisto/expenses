/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
