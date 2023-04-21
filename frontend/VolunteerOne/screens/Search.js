import React from 'react';
import { StyleSheet, Dimensions, ScrollView, TextInput } from 'react-native';
import { Block, theme } from 'galio-framework';
import UserAndOrgItem from '../components/UserAndOrgItem';
import notifications from '../constants/notifications';

const { width } = Dimensions.get('screen');

class SearchPage extends React.Component {
  state = {
    searchQuery: '',
  };

  handleSearch = query => {
    this.setState({ searchQuery: query });
  };

  renderNotifications = () => {
    const { searchQuery } = this.state;
    const filteredNotifications = notifications.filter(notification =>
      notification.title.toLowerCase().includes(searchQuery.toLowerCase()),
    );

    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
        {filteredNotifications.map(notification => (
          <UserAndOrgItem item={notification} horizontal key={notification.id} />
        ))}
      </ScrollView>
    );
  };

  render() {
    const { searchQuery } = this.state;

    return (
      <Block flex center style={styles.home}>
        <TextInput
          style={styles.searchInput}
          placeholder="Search for Volunteers or Organizations..."
          value={searchQuery}
          onChangeText={this.handleSearch}
        />
        {this.renderNotifications()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
  searchInput: {
    backgroundColor: '#fff',
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 4,
    paddingVertical: 8,
    paddingHorizontal: 12,
    marginBottom: 5,
    width: width - theme.SIZES.BASE * 2,
    marginTop: 20,
  },
});

export default SearchPage;
