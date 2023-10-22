import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddItemFormComponent } from './items/add-item-form/add-item-form.component';
import { ItemListComponent } from './items/item-list/item-list.component';
import { ItemComponent } from './items/item/item.component';

const routes: Routes = [
    { path: 'add-item', component: AddItemFormComponent },
    { path: 'list-items', component: ItemListComponent },
    { path: 'item/:id', component: ItemComponent },
    { path: '', redirectTo: '/', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
