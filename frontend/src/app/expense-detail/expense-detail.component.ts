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
        this.expense = new Expense(null, 1, user.id);
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
