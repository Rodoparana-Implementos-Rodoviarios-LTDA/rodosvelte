import type { IconDashboard, IconPhone } from '@tabler/icons-svelte';

export interface MenuItem {
	name: string;
	icon: typeof IconDashboard | typeof IconPhone;
	link: string;
	external?: boolean;
}

export interface MenuSection {
	title: string;
	link?: string;
	items: MenuItem[];
}
