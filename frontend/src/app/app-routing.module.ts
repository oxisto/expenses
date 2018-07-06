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

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { AuthGuard } from './auth.guard';
import { ExpenseDetailComponent } from './expense-detail/expense-detail.component';
import { ExpenseListComponent } from './expense-list/expense-list.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: '/expenses',
    pathMatch: 'full'
  },
  {
    path: 'expenses',
    component: ExpenseListComponent,
    canActivate: [
      AuthGuard
    ]
  },
  {
    path: 'expenses/:id',
    component: ExpenseDetailComponent,
    canActivate: [
      AuthGuard
    ]
  },
  {
    path: 'about',
    component: AboutComponent
  },
  {
    path: 'login',
    component: LoginComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
