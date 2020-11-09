import React, {Component} from 'react';
import {View, ActivityIndicator, Linking} from 'react-native';
import {StackActions} from '@react-navigation/native';
import WebView from 'react-native-webview';

class WebViewScreen extends Component {
  renderLoadingView() {
    return (
      <View style={{flex: 1, alignItems: 'center', justifyContent: 'center'}}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  render() {
    const {navigation} = this.props;
    const {url} = this.props.route.params;

    return (
      <WebView
        source={{uri: url}}
        renderLoading={this.renderLoadingView}
        startInLoadingState={true}
        ref={(ref) => {
          this.webview = ref;
        }}
        onNavigationStateChange={(event) => {
          if (event.url !== url) {
            console.log(event.url);

            this.webview.stopLoading();

            if (event.url.includes('https://www.google.com/')) {
              navigation.navigate('Internal', {url: event.url});
            } else if (event.url.includes('http://local/')) {
              const pageLink = event.url.substring(
                event.url.lastIndexOf('/') + 1,
                event.url.length,
              );

              navigation.navigate(pageLink);

              this.webview.reload();
            } else {
              Linking.openURL(event.url);
            }
          }
        }}
      />
    );
  }
}

export default WebViewScreen;
