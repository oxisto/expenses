import { Component, OnInit } from '@angular/core';
import { Expense } from '../expense';
import { ActivatedRoute, Router, ParamMap } from '@angular/router';

import { HttpClient } from '@angular/common/http';
import { ExpenseService } from '../expense.service';

@Component({
  selector: 'app-expense-detail',
  templateUrl: './expense-detail.component.html',
  styleUrls: ['./expense-detail.component.css']
})
export class ExpenseDetailComponent implements OnInit {

  expense: Expense;
  new = false;
  submitted: boolean;

  constructor(private route: ActivatedRoute,
    private router: Router,
    private http: HttpClient,
    private expenseService: ExpenseService) { }

  ngOnInit() {
    this.route.paramMap.forEach((params: ParamMap) => {
      const id = params.get('id');

      if (id === 'new') {
        this.new = true;

        this.expense = new Expense(null, 1, '1');
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
