import { Component } from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent {

    navOptions = [
        {
            title: 'Items',
            id: 'itemDropdown',
            actions: [
                { label: 'Add Item', route: '/add-item' },
                { label: 'Item List', route: '/list-items' }
            ]
        },
        {
            title: 'Allowances',
            id: 'allowanceDropdown',
            actions: [
                { label: 'Add Allowance', route: '/add-allowance' },
                { label: 'Allowance List', route: '/list-allowances' }
            ]
        },
        {
            title: 'Inventory',
            id: 'inventoryDropdown',
            actions: [
                { label: 'Create Location', route: '/add-location' },
                { label: 'Re-order Stock', route: '/order-stock' },
                { label: 'Schedule Inventory Count', route: '/schedule-inventory' },
                { label: 'Post Inventory Counts', route: '/post-inventory' },
                { label: 'Offload Excess Stock', route: '/offload' },
            ]
        },
        {
            title: 'Orders',
            id: 'ordersDropdown',
            actions: [
                { label: 'New Order', route: '/new-order' },
                { label: 'Update Order Status', route: '/order-status' },
                { label: 'Active Orders', route: '/orders' },
                { label: 'Order History', route: '/order-history' },
            ]
        }
    ]

}
