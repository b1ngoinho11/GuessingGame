import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";

export default function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [guess, setGuess] = useState("");

  const handleAuth = () => setIsAuthenticated(!isAuthenticated);
  const handleGuess = () => alert(`Your guess: ${guess}`);

  return (
    <div className="flex flex-col min-h-screen items-center justify-center bg-neutral-800 text-white dark">
      <div className="absolute top-4 right-4">
        {isAuthenticated ? (
          <DropdownMenu>
            <DropdownMenuTrigger>
              <Avatar>
                <AvatarImage src="https://github.com/shadcn.png" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
            </DropdownMenuTrigger>
            <DropdownMenuContent
              align="end"
              className="bg-card text-card-foreground"
            >
              <DropdownMenuItem onClick={handleAuth}>Logout</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button onClick={handleAuth}>Login</Button>
        )}
      </div>

      <Card className="w-full max-w-md shadow-lg bg-card text-card-foreground">
        <CardHeader>
          <h2 className="text-xl font-semibold text-center">Guessing Game</h2>
        </CardHeader>
        <CardContent className="flex flex-col gap-4">
          <Input
            type="text"
            placeholder="Enter your guess..."
            value={guess}
            onChange={(e) => setGuess(e.target.value)}
            className="bg-input text-input-foreground"
          />
        </CardContent>
        <CardFooter className="flex justify-center">
          <Button onClick={handleGuess}>Submit Guess</Button>
        </CardFooter>
      </Card>
    </div>
  );
}
