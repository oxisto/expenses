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
import { AuthService } from '../auth.service';
import { Expense } from '../expense';
import { ExpenseService } from '../expense.service';
import { User } from '../user';

@Component({
  selector: 'app-expense-detail',
  templateUrl: './expense-detail.component.html',
  styleUrls: ['./expense-detail.component.css']
})
export class ExpenseDetailComponent implements OnInit {

  expense: Expense;
  new = false;
  submitted: boolean;

  users: User[] = [];

  constructor(private route: ActivatedRoute,
    private router: Router,
    private authService: AuthService,
    private expenseService: ExpenseService) { }

  ngOnInit() {
    // for now, only push the current user, in the future, we want to support
    // access to multiple accounts
    const user = this.authService.getUser();

    this.users.push(user);

    this.route.paramMap.forEach((params: ParamMap) => {
      const id = params.get('id');

      if (id === 'new') {
        this.new = true;

        // set the account owner as default
        this.expense = new Expense(null, null, user.id, new Date());
      } else {
        this.expenseService.getExpense(id).subscribe(expense => {
          this.expense = expense;
        });
      }
    });
  }

  onSubmit() {
    this.submitted = true;

    let obs;

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
