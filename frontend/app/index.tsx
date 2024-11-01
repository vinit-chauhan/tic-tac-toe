import { Button } from "react-native";
import { AuthProvider, useAuth } from "./context/AuthContext";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import Login from "./screens/Login";
import Profile from "./screens/Profile";

const Stack = createNativeStackNavigator();

export default function Root() {
  return (
    <AuthProvider>
      <Layout />
    </AuthProvider>
  );
}

export const Layout = () => {
  const { authState, onLogout } = useAuth();

  return (
    <Stack.Navigator>
      {authState?.authenticated ? (
        <Stack.Screen
          name="profile"
          component={Profile}
          options={{
            headerRight: () => <Button onPress={onLogout} title="Sign out" />,
          }}
        />
      ) : (
        <Stack.Screen name="Login" component={Login} />
      )}
    </Stack.Navigator>
  );
};
