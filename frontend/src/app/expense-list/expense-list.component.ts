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

import { AfterViewInit, Component, OnInit } from '@angular/core';
import { AccountService } from '../account.service';
import { Expense } from '../expense';
import { ExpenseService } from '../expense.service';
import { User } from '../user';

@Component({
  selector: 'app-expense-list',
  templateUrl: './expense-list.component.html',
  styleUrls: ['./expense-list.component.css']
})
export class ExpenseListComponent implements OnInit, AfterViewInit {

  users: User[];
  expenses: Expense[];

  constructor(private expenseService: ExpenseService,
    private accountService: AccountService) { }

  ngOnInit() {
    this.accountService.getAccounts().subscribe(users => this.users = users);

    this.updateExpenses();
  }

  ngAfterViewInit() {
    this.expenseService.syncOfflineExpenses().subscribe(() => {
      this.updateExpenses();
    });
  }

  updateExpenses() {
    this.expenseService.getExpenses()
      .subscribe(expenses => {
        this.expenses = expenses;
      });
  }

}
