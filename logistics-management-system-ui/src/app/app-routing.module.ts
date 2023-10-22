import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddItemFormComponent } from './add-item-form/add-item-form.component';

const routes: Routes = [
    { path: 'add-item', component: AddItemFormComponent },
    { path: '', redirectTo: '/home', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
