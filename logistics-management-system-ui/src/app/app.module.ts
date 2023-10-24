import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AddItemFormComponent } from './items/add-item-form/add-item-form.component';
import { ItemListComponent } from './items/item-list/item-list.component';
import { HttpClientModule } from '@angular/common/http';
import { ItemComponent } from './items/item/item.component';
import { AddAllowanceFormComponent } from './allowances/add-allowance-form/add-allowance-form.component';
import { AllowanceListComponent } from './allowances/allowance-list/allowance-list.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { InventoryValueComponent } from './dashboard/widget/inventory-value/inventory-value.component';
import { OrderVolumeComponent } from './dashboard/widget/order-volume/order-volume.component';
import { IssueEffectivenessComponent } from './dashboard/widget/issue-effectiveness/issue-effectiveness.component';
import { OpenOrdersComponent } from './dashboard/widget/open-orders/open-orders.component';
import { StockLevelsComponent } from './dashboard/widget/stock-levels/stock-levels.component';
import { NavbarComponent } from './navbar/navbar.component';

@NgModule({
  declarations: [
    AppComponent,
    AddItemFormComponent,
    ItemListComponent,
    ItemComponent,
    AddAllowanceFormComponent,
    AllowanceListComponent,
    DashboardComponent,
    InventoryValueComponent,
    OrderVolumeComponent,
    IssueEffectivenessComponent,
    OpenOrdersComponent,
    StockLevelsComponent,
    NavbarComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    NgbModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
