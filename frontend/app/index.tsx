import { Button, StyleSheet, View } from "react-native";
import { AuthProvider, useAuth } from "./context/AuthContext";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import Login from "./screens/Login";
import Signup from "./screens/Signup";
import Profile from "./screens/Profile";
import { RootStackParamList } from "./navigarionTypes";

const Stack = createNativeStackNavigator<RootStackParamList>();

const linking = {
  prefixes: ["yourapp://", "https://yourapp.com"],
  config: {
    screens: {
      Login: "login",
      Signup: "signup",
      Profile: "profile",
    },
  },
};

export default function App() {
  return (
    <AuthProvider>
      <Layout />
    </AuthProvider>
  );
}

const Layout = () => {
  const { authState, onLogout } = useAuth();

  return (
    <Stack.Navigator>
      {authState?.authenticated ? (
        <Stack.Screen
          name="Profile"
          component={Profile}
          options={{
            headerRight: () => (
              <View style={{ marginRight: 10 }}>
                <Button title="Logout" onPress={onLogout} />{" "}
              </View>
            ),
          }}
        />
      ) : (
        <>
          <Stack.Screen name="Login" component={Login} />
          <Stack.Screen
            name="Signup"
            options={{ title: "Sign up" }}
            component={Signup}
          />
        </>
      )}
    </Stack.Navigator>
  );
};
