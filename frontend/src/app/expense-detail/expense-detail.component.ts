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

import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap, Router } from '@angular/router';
import { NgbDateAdapter, NgbDateNativeAdapter, NgbTimeStruct } from '../../../node_modules/@ng-bootstrap/ng-bootstrap';
import { AccountService } from '../account.service';
import { AuthService } from '../auth.service';
import { Expense } from '../expense';
import { ExpenseService } from '../expense.service';
import { User } from '../user';

@Component({
  selector: 'app-expense-detail',
  templateUrl: './expense-detail.component.html',
  styleUrls: ['./expense-detail.component.css'],
  providers: [{ provide: NgbDateAdapter, useClass: NgbDateNativeAdapter }]
})
export class ExpenseDetailComponent implements OnInit {
  // need to split day and time for our date and time picker
  date: Date;
  time: NgbTimeStruct;

  expense: Expense;
  new = false;
  submitted: boolean;

  users: User[] = [];

  constructor(private route: ActivatedRoute,
    private router: Router,
    private authService: AuthService,
    private accountService: AccountService,
    private expenseService: ExpenseService) { }

  ngOnInit() {
    this.accountService.getAccounts().subscribe(users => this.users = Object.values(users));

    // default user
    const user = this.authService.getUser();

    this.route.paramMap.forEach((params: ParamMap) => {
      const id = params.get('id');

      if (id === 'new') {
        this.new = true;

        // set the account owner as default
        this.expense = new Expense(null, null, user.id, new Date());

        // prepare date
        this.splitDate();
      } else {
        this.expenseService.getExpense(id).subscribe(expense => {
          this.expense = expense;

          // prepare date
          this.splitDate();
        });
      }
    });
  }

  splitDate() {
    this.date = new Date(this.expense.timestamp);
    this.time = {
      minute: this.date.getMinutes(),
      hour: this.date.getHours(),
      second: this.date.getSeconds()
    };
  }

  mergeDate() {
    this.expense.timestamp = this.date;
    this.expense.timestamp.setSeconds(this.time.second);
    this.expense.timestamp.setMinutes(this.time.minute);
    this.expense.timestamp.setHours(this.time.hour);
  }

  onDelete() {
    this.expenseService.deleteExpense(this.expense.id).subscribe(() => {
      this.router.navigateByUrl('/');
    });
  }

  onSubmit() {
    this.submitted = true;

    let obs;

    // merge date
    this.mergeDate();

    if (this.expense.id == null) {
      obs = this.expenseService.postExpense(this.expense);
    } else {
      obs = this.expenseService.putExpense(this.expense);
    }

    obs.subscribe(() => {
      this.router.navigateByUrl('/');
    });
  }

}
