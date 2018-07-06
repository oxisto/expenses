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

import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable, ObservableInput, of } from 'rxjs';
import { catchError, startWith } from 'rxjs/operators';
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
        catchError(this.authService.handleHttpError.bind(this.authService)),
        startWith(this.getExpenseCache())
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
        catchError((err: HttpErrorResponse): ObservableInput<any> => {
          if (err.status === 504) {
            // try to store it offline
            return this.saveExpenseOffline(expense);
          }

          return EMPTY;
        })
      );
  }

  putExpense(expense: Expense): Observable<Expense> {
    return this.http.put<Expense>(EXPENSE_ENDPOINT + '/' + expense.id, expense)
      .pipe(
        catchError(this.authService.handleHttpError.bind(this.authService))
      );
  }

  getExpenseCache(): Array<Expense> {
    let cache = JSON.parse(localStorage.getItem('expenseCache'));

    if (cache === null) {
      cache = new Array();
    }

    return cache;
  }

  setExpenseCache(cache: Array<Expense>) {
    localStorage.setItem('expenseCache', JSON.stringify(cache));
  }

  saveExpenseOffline(expense: Expense): Observable<Expense> {
    const cache = this.getExpenseCache();

    expense.localStorage = true;
    cache.push(expense);

    this.setExpenseCache(cache);

    return of(expense);
  }

  removeOfflineExpense(index: number) {
    const cache = this.getExpenseCache();

    if (index < cache.length) {
      cache.splice(index, 1);
    }

    this.setExpenseCache(cache);
  }

  syncOfflineExpenses(): Observable<Expense[]> {
    let emitter;
    const obs = Observable.create(e => emitter = e);
    const cache = this.getExpenseCache();

    cache.forEach((element, index) => {
      // remove it from the cache, to avoid storing it twice if upload fails again
      this.removeOfflineExpense(index);

      const expense = cache[index];

      // try to post it again, if it fails it will go back into the cache
      this.postExpense(expense).subscribe(e => {
        if (!e.localStorage) {
          emitter.next(e);
        }
      });
    });

    return obs;
  }
}
