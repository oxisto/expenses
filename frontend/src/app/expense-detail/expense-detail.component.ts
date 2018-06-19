import { Component, OnInit } from '@angular/core';
import { Expense } from '../expense';
import { ActivatedRoute, Router, ParamMap } from '@angular/router';

import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-expense-detail',
  templateUrl: './expense-detail.component.html',
  styleUrls: ['./expense-detail.component.css']
})
export class ExpenseDetailComponent implements OnInit {

  expense: Expense;
  new = false;
  submitted: boolean;

  constructor(private route: ActivatedRoute, private router: Router, private http: HttpClient) { }

  ngOnInit() {
    this.route.paramMap.forEach((params: ParamMap) => {
      const id = params.get('id');

      if (id === 'new') {
        this.new = true;

        this.expense = new Expense(1, 1);
      }
    });
  }

  onSubmit() {
    this.submitted = true;

    this.http.post('/api/expenses', this.expense)
      .subscribe(() => {
        this.router.navigateByUrl('/');
      });
  }

}
