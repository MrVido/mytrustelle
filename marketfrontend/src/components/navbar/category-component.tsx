import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuIndicator,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  NavigationMenuViewport,
} from "@/components/ui/navigation-menu";
import { cn } from "@/lib/utils";
import React from "react";

export const categoryComponent: { title: string; href: string; description: string }[] = [
  {
    title: "Cars and Vehicles",
    href: "/carsandvehicles",
    description:
      "The Cars And vehicles For Sale from Private Sellers and Dealerships for MyTrustElle",
  },
  {
    title: "Housing",
    href: "/housing",
    description:
      "The Housing Section provides rentals, homes for sale, condos from Private Sellers and Real Estate Agents and agencys for MyTrustElle",
  },
  {
    title: "Jobs",
    href: "/job",
    description:
      "The job section if you're looking for your next opportunity from Private Companies, Public companies and agencys for MyTrustElle",
  },
];

export const ListItem = React.forwardRef<
  React.ElementRef<"a">,
  React.ComponentPropsWithoutRef<"a">
>(({ className, title, children, ...props }, ref) => {
  return (
    <li>
      <NavigationMenuLink asChild>
        <a
          ref={ref}
          className={cn(
            "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
            className
          )}
          {...props}
        >
          <div className="text-sm font-medium leading-none">{title}</div>
          <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
            {children}
          </p>
        </a>
      </NavigationMenuLink>
    </li>
  );
});
ListItem.displayName = "ListItem";
