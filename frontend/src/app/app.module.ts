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

import { HttpClientModule } from '@angular/common/http';
import { APP_INITIALIZER, NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { ServiceWorkerModule } from '@angular/service-worker';
import { JwtModule } from '@auth0/angular-jwt';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { MomentModule } from 'ngx-moment';
import { environment } from '../environments/environment';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { getToken } from './auth.service';
import { ConfigService } from './config.service';
import { ExpenseDetailComponent } from './expense-detail/expense-detail.component';
import { ExpenseListComponent } from './expense-list/expense-list.component';
import { LoginComponent } from './login/login.component';
import { AboutComponent } from './about/about.component';

@NgModule({
  declarations: [
    AppComponent,
    ExpenseListComponent,
    ExpenseDetailComponent,
    LoginComponent,
    AboutComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    MomentModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: getToken,
        blacklistedRoutes: ['/auth/']
      }
    }),
    NgbModule.forRoot(),
    ServiceWorkerModule.register('/ngsw-worker.js', { enabled: environment.production })
  ],
  providers: [
    {
      provide: APP_INITIALIZER,
      useFactory: (config: ConfigService) => () => config.load(),
      deps: [ConfigService], multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
