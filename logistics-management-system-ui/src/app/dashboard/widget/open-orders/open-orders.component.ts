import { Component } from '@angular/core';

@Component({
  selector: 'app-open-orders',
  templateUrl: './open-orders.component.html',
  styleUrls: ['./open-orders.component.scss']
})
export class OpenOrdersComponent {

    orders = [
        {id: "00000001", itemId: "987654321", itemName: "item name", quantity: 0, orderDate: "10/22/2023", status: "created",  statusDate: "10/22/2023"},
        {id: "00000001", itemId: "987654321", itemName: "item name", quantity: 0, orderDate: "10/22/2023", status: "created",  statusDate: "10/22/2023"},
        {id: "00000001", itemId: "987654321", itemName: "item name", quantity: 0, orderDate: "10/22/2023", status: "created",  statusDate: "10/22/2023"},
        {id: "00000001", itemId: "987654321", itemName: "item name", quantity: 0, orderDate: "10/22/2023", status: "created",  statusDate: "10/22/2023"},
        {id: "00000001", itemId: "987654321", itemName: "item name", quantity: 0, orderDate: "10/22/2023", status: "created",  statusDate: "10/22/2023"},
        {id: "00000001", itemId: "987654321", itemName: "item name", quantity: 0, orderDate: "10/22/2023", status: "created",  statusDate: "10/22/2023"},
    ]
}
