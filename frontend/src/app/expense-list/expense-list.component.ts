import { Component, OnInit } from '@angular/core';
import { Expense } from '../expense';
import { ExpenseService } from '../expense.service';

@Component({
  selector: 'app-expense-list',
  templateUrl: './expense-list.component.html',
  styleUrls: ['./expense-list.component.css']
})
export class ExpenseListComponent implements OnInit {

  expenses: Expense[];

  constructor(private expenseService: ExpenseService) { }

  ngOnInit() {
    this.expenseService.getExpenses().subscribe(expenses => {
      this.expenses = expenses;
    });
  }

}
